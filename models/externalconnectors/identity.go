// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0 "github.com/aurva-io/avmsgraph-sdk-go/models"
)

type Identity struct {
    iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.Entity
}
// NewIdentity instantiates a new Identity and sets the default values.
func NewIdentity()(*Identity) {
    m := &Identity{
        Entity: *iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.NewEntity(),
    }
    return m
}
// CreateIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Identity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*IdentityType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type of identity. Possible values are: user or group for Microsoft Entra identities and externalgroup for groups in an external system.
// returns a *IdentityType when successful
func (m *Identity) GetTypeEscaped()(*IdentityType) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IdentityType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Identity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeEscaped sets the type property value. The type of identity. Possible values are: user or group for Microsoft Entra identities and externalgroup for groups in an external system.
func (m *Identity) SetTypeEscaped(value *IdentityType)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
type Identityable interface {
    iff6ba9ca3c965d94e15ab7aa30386d60bd694ee3b15d5295d3980a9418fbbab0.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*IdentityType)
    SetTypeEscaped(value *IdentityType)()
}
