// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0 "github.com/aurva-io/avmsgraph-sdk-go/models"
)

type ExternalActivityResult struct {
    ExternalActivity
}
// NewExternalActivityResult instantiates a new ExternalActivityResult and sets the default values.
func NewExternalActivityResult()(*ExternalActivityResult) {
    m := &ExternalActivityResult{
        ExternalActivity: *NewExternalActivity(),
    }
    return m
}
// CreateExternalActivityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalActivityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalActivityResult(), nil
}
// GetError gets the error property value. Error information that explains the failure to process an external activity.
// returns a PublicErrorable when successful
func (m *ExternalActivityResult) GetError()(iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable) {
    val, err := m.GetBackingStore().Get("error")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalActivityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalActivity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalActivityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error information that explains the failure to process an external activity.
func (m *ExternalActivityResult) SetError(value iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable)() {
    err := m.GetBackingStore().Set("error", value)
    if err != nil {
        panic(err)
    }
}
type ExternalActivityResultable interface {
    ExternalActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable)
    SetError(value iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.PublicErrorable)()
}
