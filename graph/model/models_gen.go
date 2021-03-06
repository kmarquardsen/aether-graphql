// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SimCard     *SimCard `json:"simCard"`
}

type DeviceGroup struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Devices []*Device `json:"devices"`
	Slice   *Slice    `json:"slice"`
}

type Enterprise struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Sites        []*Site        `json:"sites"`
	Applications []*Application `json:"applications"`
}

type SimCard struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Site struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Slices      []*Slice     `json:"slices"`
	SmallCells  []*SmallCell `json:"smallCells"`
	Devices     []*Device    `json:"devices"`
	SimCards    []*SimCard   `json:"simCards"`
}

type Slice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SmallCell struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
