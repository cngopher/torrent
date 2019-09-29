package torrent

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cngopher/torrent/internal/testutil"
	"github.com/cngopher/torrent/storage"
	"github.com/stretchr/testify/require"
)

func TestHashPieceAfterStorageClosed(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	tt := &Torrent{
		storageOpener: storage.NewClient(storage.NewFile(td)),
	}
	mi := testutil.GreetingMetaInfo()
	info, err := mi.UnmarshalInfo()
	require.NoError(t, err)
	require.NoError(t, tt.setInfo(&info))
	require.NoError(t, tt.storage.Close())
	tt.hashPiece(0)
}
