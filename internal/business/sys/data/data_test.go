package data_test

import (
	"github.com/circutor/ps-go-client/internal/business/sys/data"
	models "github.com/circutor/ps-go-client/pkg/models"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBodyDecode(t *testing.T) {
	xmlFile, err := os.Open("../../sampleEntities/allDevices.xml")
	require.NoError(t, err)

	byteValue, err := io.ReadAll(xmlFile)
	require.NoError(t, err)

	decode, err := data.BodyDecode(byteValue, &models.Devices{})

	assert.NotEmpty(t, decode)
	assert.Nil(t, err)
}

func TestBodyDecodeError(t *testing.T) {
	decode, err := data.BodyDecode([]byte(""), &models.Devices{})

	assert.Empty(t, decode)
	assert.Error(t, err)
}
