package data_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/circutor/ps-go-client/internal/business/sys/data"
	models "github.com/circutor/ps-go-client/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBodyDecode(t *testing.T) {
	t.Parallel()

	xmlFile, err := os.Open("../../sampleEntities/allDevices.xml")
	require.NoError(t, err)

	byteValue, err := io.ReadAll(xmlFile)
	require.NoError(t, err)

	decode, err := data.BodyDecode(byteValue, &models.Devices{})

	assert.NotEmpty(t, decode)
	assert.Nil(t, err)
}

func TestBodyDecodeError(t *testing.T) {
	t.Parallel()

	decode, err := data.BodyDecode([]byte(""), &models.Devices{})

	assert.Empty(t, decode)
	assert.Equal(t, true, errors.Is(err, io.EOF))
}
