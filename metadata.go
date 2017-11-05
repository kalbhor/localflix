package main

import "fmt"

type MediaFile struct {
	Path        string
	Title       string
	Length      int
	TotalLength int
	Cast        []string
	Directors   []string
	Desc        string
	Rating      float32
	Uploaded    bool
	IsManual    bool
}

type Movie MediaFile
type Episode MediaFile

type Season struct {
	Title    string     // Path of the season folder
	Episodes []*Episode // Episode files in the folder
	IsManual bool
}

type Series struct {
	Title    string // Path of TV series folder (containing season folders)
	Seasons  []*Season
	IsManual bool
}

func NewMovie(title string) *Movie {
	m := &Movie{Title: title}
	return m
}

func NewSeries(title string) *Series {
	s := &Series{Title: title}
	return s
}

func NewSeason() *Season {
	s := &Season{}
	return s
}

func NewEpisode() *Episode {
	e := &Episode{}
	return e
}

func CheckSeries(s *Series) bool {
	if cap(s.Seasons) != 0 {
		return false
	}
	return true
}

func CheckSeasons(s *Season) bool {
	if cap(s.Episodes) != 0 {
		return false
	}
	return true
}

func (s *Series) AddSeason(season *Season) {
	if CheckSeries(s) {
		s.Seasons = make([]*Season, 1)
		s.Seasons[0] = season
	} else {
		s.Seasons = append(s.Seasons, season)
	}
}

func (s *Season) AddEpisode(episode *Episode) {
	if CheckSeasons(s) {
		s.Episodes = make([]*Episode, 1)
		s.Episodes[0] = episode
	} else {
		s.Episodes = append(s.Episodes, episode)
	}
}

func (s *Series) DisplayTree() {
	fmt.Println(s.Title)
	for _, season := range s.Seasons {
		fmt.Println("	", season.Title)
		for _, episodes := range season.Episodes {
			fmt.Println("		", episodes.Title)
			fmt.Println("		", episodes.Path)
		}
	}
}

func main() {
	x := NewSeries("Mr Robot")

	y := NewSeason()
	y.Title = "Season 1"
	e := NewEpisode()
	e.Title = "Pilot"
	y.AddEpisode(e)
	x.AddSeason(y)

	x.DisplayTree()

}
