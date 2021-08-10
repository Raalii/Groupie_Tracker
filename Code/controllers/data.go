package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	Model "../models"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

var ArtistsFull []Model.ArtistMain
var Artists []Model.Artist
var AllDates Model.Dates
var AllLocations Model.Locations
var AllRelations Model.Relations

func GetArtistsFromAPI() error {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

var Nationnality []string = []string{"Britannique", "Américain", "Allemand", "Américain", "Américain", "Américain", "Américain", "Américain", "Australien", "Américain", "Américaine", "Américaine", "Britannique", "Britannique", "Britannique", "Américain", "Britannique", "Britannique", "Américain", "Britannique", "Brésilien", "Américain", "Américain", "Américain", "Américain", "Américain", "Américain", "Néerlandais", "Américain", "Américain", "Américain", "Canadien", "Américain", "Américain", "Américain", "Irlandais", "Britannique", "Américain", "Britannique", "Américain", "Américain", "Américain", "Américain", "Américain", "Américain", "Britannique", "Américain", "Américain", "Britannique", "Britannique", "Américain", "Américain"}

func GetDatesFromAPI() error {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &AllDates)
	return nil
}

func GetLocationsFromAPI() error {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &AllLocations)
	return nil
}

func GetRelationsFromAPI() error {
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return errors.New("error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by ReadAll")
	}
	json.Unmarshal(bytes, &AllRelations)
	return nil
}

// Fonction qui tranforme le nom des villes pays en une syntaxe plus belle

func SpacesBetweenLess(word string) string {
	var newWord string
	for _, v := range word {
		if v == '-' {
			newWord += " - "
		} else {
			newWord += string(v)
		}
	}
	return newWord
}

func TextTransform(word string) string {
	var newWord string
	for _, v := range word {
		if v == rune(95) { // "_"
			newWord += " "
		} else {
			newWord += string(v)
		}
	}
	return newWord
}

func TextAllTransform(slice []string) []string {
	NewSlice := []string{}
	for _, v := range slice {
		NewSlice = append(NewSlice, SpacesBetweenLess(TextTransform(v)))
	}

	return NewSlice
}

func MapTransformText(maps map[string][]string) map[string][]string {
	NewMap := make(map[string][]string)
	for key, value := range maps {
		newKey := SpacesBetweenLess(TextTransform(key))
		NewMap[newKey] = value
	}
	return NewMap
}

func GetAll() error {
	if len(ArtistsFull) != 0 {
		return nil
	}
	err1 := GetArtistsFromAPI()
	err2 := GetLocationsFromAPI()
	err3 := GetDatesFromAPI()
	err4 := GetRelationsFromAPI()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return errors.New("error by get data artists, locations, dates")
	}

	for i := range Artists {
		var tmpl Model.ArtistMain
		tmpl.ID = i + 1
		tmpl.Image = Artists[i].Image
		tmpl.Name = Artists[i].Name
		tmpl.Members = Artists[i].Members
		tmpl.CreationDate = Artists[i].CreationDate
		tmpl.FirstAlbum = Artists[i].FirstAlbum
		tmpl.Locations = AllLocations.Index[i].Locations
		tmpl.ConcertDates = AllDates.Index[i].Dates
		tmpl.DatesLocations = MapTransformText(AllRelations.Index[i].DatesLocations)
		tmpl.Country = Nationnality[i]
		if len(tmpl.Members) > 1 {
			tmpl.Type = "Groupe"
		} else {
			tmpl.Type = "Artiste"
		}
		ArtistsFull = append(ArtistsFull, tmpl)
	}
	return nil
}

// array := []string{"boston-usa", "chicago-usa", "north_carolina-usa","georgia-usa","los_angeles-usa","saitama-japan","osaka-japan","nagoya-japan","penrose-new_zealand","dunedin-new_zealand"}
