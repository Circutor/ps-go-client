package data_test

import (
	"io"
	"os"
	"ps-go-client/internal/business/model"
	"ps-go-client/internal/business/sys/data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBodyDecode(t *testing.T) {
	xmlFile, err := os.Open("../../sampleEntities/allDevices.xml")
	require.NoError(t, err)

	byteValue, err := io.ReadAll(xmlFile)
	require.NoError(t, err)

	decode, err := data.BodyDecode(byteValue, &model.Devices{})

	assert.NotEmpty(t, decode)
	assert.Nil(t, err)
}

func TestBodyDecodeError(t *testing.T) {
	decode, err := data.BodyDecode([]byte(""), &model.Devices{})

	assert.Empty(t, decode)
	assert.Error(t, err)
}
