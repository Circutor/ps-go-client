package data_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"ps-go-client/internal/business/model"
	"ps-go-client/internal/business/sys/data"
	"testing"
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
