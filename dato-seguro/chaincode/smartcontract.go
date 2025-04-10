// SPDX-License-Identifier: Apache-2.0
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Licencia struct {
	Numero       string `json:"numero"`
	FechaEmision string `json:"fechaEmision"`
	Tipo         string `json:"tipo"`
	Estado       string `json:"estado"`
}

type Infraccion struct {
	Codigo string `json:"codigo"`
	Fecha  string `json:"fecha"`
	Puntos int    `json:"puntos"`
}

type Ciudadano struct {
	// Registro Civil
	Cedula                string `json:"cedula"`
	NombresCompletos      string `json:"nombresCompletos"`
	Sexo                  string `json:"sexo"`
	CondicionCiudadano    string `json:"condicionCiudadano"`
	FechaNacimiento       string `json:"fechaNacimiento"`
	LugarNacimiento       string `json:"lugarNacimiento"`
	Nacionalidad          string `json:"nacionalidad"`
	EstadoCivil           string `json:"estadoCivil"`
	CodigoDactilar        string `json:"codigoDactilar"`
	Conyuge               string `json:"conyuge"`
	NombrePadre           string `json:"nombrePadre"`
	NacionalidadPadre     string `json:"nacionalidadPadre"`
	NombreMadre           string `json:"nombreMadre"`
	NacionalidadMadre     string `json:"nacionalidadMadre"`
	Domicilio             string `json:"domicilio"`
	CallesDomicilio       string `json:"callesDomicilio"`
	NumeroCasa            string `json:"numeroCasa"`
	FechaMatrimonio       string `json:"fechaMatrimonio"`
	LugarMatrimonio       string `json:"lugarMatrimonio"`
	FechaDefuncion        string `json:"fechaDefuncion"`
	Observaciones         string `json:"observaciones"`
	FechaInscripcionDef   string `json:"fechaInscripcionDefuncion"`
	FechaExpedicionCedula string `json:"fechaExpedicionCedula"`
	FechaExpiracionCedula string `json:"fechaExpiracionCedula"`
	Genero                string `json:"genero"`

	// Consejo Nacional Electoral
	Provincia       string `json:"provincia"`
	Canton          string `json:"canton"`
	Parroquia       string `json:"parroquia"`
	Recinto         string `json:"recinto"`
	CodigoElectoral string `json:"codigoElectoral"`

	// Agencia Nacional de Tránsito
	Licencias    []Licencia   `json:"licencias"`
	Infracciones []Infraccion `json:"infracciones"`

	DocType string `json:"docType"`
}

func (s *SmartContract) CrearCiudadano(ctx contractapi.TransactionContextInterface, ciudadanoJSON string) error {
	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}
	if clientMSP != "RegistroCivilMSP" {
		return fmt.Errorf("solo Registro Civil puede crear ciudadanos")
	}
	var ciudadano Ciudadano
	err = json.Unmarshal([]byte(ciudadanoJSON), &ciudadano)
	if err != nil {
		return err
	}
	ciudadano.DocType = "ciudadano"
	ciudadanoKey := "CIUDADANO_" + ciudadano.Cedula
	exists, err := ctx.GetStub().GetState(ciudadanoKey)
	if err != nil {
		return err
	}
	if exists != nil {
		return fmt.Errorf("ciudadano con cédula %s ya existe", ciudadano.Cedula)
	}

	ciudadanoBytes, err := json.Marshal(ciudadano)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(ciudadanoKey, ciudadanoBytes)
}

func (s *SmartContract) LeerCiudadano(ctx contractapi.TransactionContextInterface, cedula string) (*Ciudadano, error) {
	ciudadanoKey := "CIUDADANO_" + cedula
	data, err := ctx.GetStub().GetState(ciudadanoKey)
	if err != nil || data == nil {
		return nil, fmt.Errorf("ciudadano no encontrado")
	}
	var ciudadano Ciudadano
	err = json.Unmarshal(data, &ciudadano)
	return &ciudadano, err
}

func (s *SmartContract) CiudadanoExiste(ctx contractapi.TransactionContextInterface, cedula string) (bool, error) {
	ciudadanoKey := "CIUDADANO_" + cedula
	data, err := ctx.GetStub().GetState(ciudadanoKey)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func (s *SmartContract) ActualizarDatosElectorales(ctx contractapi.TransactionContextInterface, cedula string, provincia, canton, parroquia, recinto, codigoElectoral string) error {
	clientMSP, _ := ctx.GetClientIdentity().GetMSPID()
	if clientMSP != "CneMSP" {
		return fmt.Errorf("solo CNE puede actualizar datos electorales")
	}
	ciudadano, err := s.LeerCiudadano(ctx, cedula)
	if err != nil {
		return err
	}
	ciudadano.Provincia = provincia
	ciudadano.Canton = canton
	ciudadano.Parroquia = parroquia
	ciudadano.Recinto = recinto
	ciudadano.CodigoElectoral = codigoElectoral
	ciudadanoBytes, _ := json.Marshal(ciudadano)
	return ctx.GetStub().PutState("CIUDADANO_"+cedula, ciudadanoBytes)
}

func (s *SmartContract) AgregarLicencia(ctx contractapi.TransactionContextInterface, cedula string, numero, fechaEmision, tipo, estado string) error {
	clientMSP, _ := ctx.GetClientIdentity().GetMSPID()
	if clientMSP != "AntMSP" {
		return fmt.Errorf("solo ANT puede agregar licencias")
	}
	ciudadano, err := s.LeerCiudadano(ctx, cedula)
	if err != nil {
		return err
	}
	lic := Licencia{
		Numero: numero, FechaEmision: fechaEmision, Tipo: tipo, Estado: estado,
	}
	ciudadano.Licencias = append(ciudadano.Licencias, lic)
	ciudadanoBytes, _ := json.Marshal(ciudadano)
	return ctx.GetStub().PutState("CIUDADANO_"+cedula, ciudadanoBytes)
}

func (s *SmartContract) AgregarInfraccion(ctx contractapi.TransactionContextInterface, cedula string, codigo, fecha string, puntos int) error {
	clientMSP, _ := ctx.GetClientIdentity().GetMSPID()
	if clientMSP != "AntMSP" {
		return fmt.Errorf("solo ANT puede agregar infracciones")
	}
	ciudadano, err := s.LeerCiudadano(ctx, cedula)
	if err != nil {
		return err
	}
	inf := Infraccion{Codigo: codigo, Fecha: fecha, Puntos: puntos}
	ciudadano.Infracciones = append(ciudadano.Infracciones, inf)
	ciudadanoBytes, _ := json.Marshal(ciudadano)
	return ctx.GetStub().PutState("CIUDADANO_"+cedula, ciudadanoBytes)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err.Error())
	}
	if err := chaincode.Start(); err != nil {
		panic(err.Error())
	}
}
