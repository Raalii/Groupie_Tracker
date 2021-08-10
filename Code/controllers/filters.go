package controllers

import (
	"fmt"
	"time"

	Model "../models"
)

func FilterArtistsByName(pattern string, Artists []Model.ArtistMain) []Model.ArtistMain {
	if pattern == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	for _, Artist := range Artists {
		if Artist.Name == pattern {
			ArtistsFiltered = append(ArtistsFiltered, Artist)
			return ArtistsFiltered
		}

		for _, Member := range Artist.Members {
			if Member == pattern {
				ArtistsFiltered = append(ArtistsFiltered, Artist)
				return ArtistsFiltered
			}
		}
	}

	return ArtistsFiltered
}

func FilterArtistsByCreationDateOfAlbum(start, end string, Artists []Model.ArtistMain) []Model.ArtistMain {
	if start == "" || end == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	dateStart, err1 := time.Parse("2006-01-02", start)
	dateEnd, err2 := time.Parse("2006-01-02", end)

	if err1 != nil || err2 != nil {
		fmt.Println("ERROR 1 : ", err1)
		fmt.Println("ERROR 2 : ", err2)
		return Artists
	}

	for _, value := range Artists {
		CurrArtistDate, err := time.Parse("02-01-2006", value.FirstAlbum)
		if err != nil {
			fmt.Println("ERROR : ", err)
			return []Model.ArtistMain{}
		}

		if CurrArtistDate.After(dateStart) && CurrArtistDate.Before(dateEnd) {
			ArtistsFiltered = append(ArtistsFiltered, value)
		}

	}
	return ArtistsFiltered
}

func FilterArtistsByCreationDate(start, end string, Artists []Model.ArtistMain) []Model.ArtistMain {
	if start == "" || end == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	dateStart, err1 := time.Parse("2006-01-02", start)
	dateEnd, err2 := time.Parse("2006-01-02", end)

	YearStart := dateStart.Year()
	YearEnd := dateEnd.Year()

	if err1 != nil || err2 != nil {
		fmt.Println("ERROR 1 : ", err1)
		fmt.Println("ERROR 2 : ", err2)
		return Artists
	}

	for _, value := range Artists {
		if YearStart <= value.CreationDate && value.CreationDate <= YearEnd {
			ArtistsFiltered = append(ArtistsFiltered, value)
		}
	}
	return ArtistsFiltered
}

func FiltersConcertByName(pattern string, Artists []Model.ArtistMain) []Model.ArtistMain {
	if pattern == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	for _, v := range Artists {
		if pattern == v.Name {
			ArtistsFiltered = append(ArtistsFiltered, v)
			return ArtistsFiltered
		}

		for _, Members := range v.Members {
			if pattern == Members {
				ArtistsFiltered = append(ArtistsFiltered, v)
				return ArtistsFiltered
			}
		}
	}
	return ArtistsFiltered
}

func FilterArtistsByType(Types string, Artists []Model.ArtistMain) []Model.ArtistMain {
	ArtistsFiltered := []Model.ArtistMain{}

	for _, Artist := range Artists {
		if Artist.Type == Types {
			ArtistsFiltered = append(ArtistsFiltered, Artist)
		}
	}

	return ArtistsFiltered
}

func DisplayArtistsByTypesWithForm(GroupForm string, ArtistForm string, Artists []Model.ArtistMain) []Model.ArtistMain {
	if GroupForm == "true" && ArtistForm == "true" {
		return Artists
	} else if GroupForm == "true" {
		return FilterArtistsByType("Groupe", Artists)
	} else if ArtistForm == "true" {
		return FilterArtistsByType("Artiste", Artists)
	} else {
		return Artists
	}
}

func FilterArtistByCountry(Country string, Artists []Model.ArtistMain) []Model.ArtistMain {

	if Country == "all" || Country == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	for _, Artist := range Artists {
		if Artist.Country == Country {
			ArtistsFiltered = append(ArtistsFiltered, Artist)
		}
	}

	return ArtistsFiltered
}

func FilterConcertByLocations(pattern string, Artists []Model.ArtistMain) []Model.ArtistMain {

	if pattern == "" {
		return Artists
	}

	ArtistsFiltered := []Model.ArtistMain{}

	for _, Artist := range Artists {
		Artist.DatesLocations = FilterLocation(pattern, Artist.DatesLocations)
		ArtistsFiltered = append(ArtistsFiltered, Artist)
	}

	return ArtistsFiltered
}

func FilterLocation(pattern string, maps map[string][]string) map[string][]string {
	NewMap := make(map[string][]string)
	for Location, Date := range maps {
		if Location == pattern {
			NewMap[Location] = Date
		}
	}
	return NewMap
}

func ConcertFilterSearchAll(pattern string, Artists []Model.ArtistMain) []Model.ArtistMain {
	ResultSearchLocation := FilterConcertByLocations(pattern, Artists)
	ResultSearchName := FiltersConcertByName(pattern, Artists)
	if len(ResultSearchLocation) == 0 && len(ResultSearchName) == 0 {
		return []Model.ArtistMain{}
	} else if len(ResultSearchLocation) > 0 && len(ResultSearchName) == 0 {
		return ResultSearchLocation
	} else {
		return ResultSearchName
	}
}
