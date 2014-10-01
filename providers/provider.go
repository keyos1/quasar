package providers

import (
	"github.com/steeve/pulsar/bittorrent"
	"github.com/steeve/pulsar/tmdb"
	"github.com/steeve/pulsar/trakt"
)

type Searcher interface {
	SearchLinks(query string) []*bittorrent.Torrent
}

type MovieSearcher interface {
	SearchMovieLinks(movie *tmdb.Movie) []*bittorrent.Torrent
}

type EpisodeSearcher interface {
	SearchEpisodeLinks(episode *trakt.ShowEpisode) []*bittorrent.Torrent
}