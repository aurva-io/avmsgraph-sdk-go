// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnAttributeCollectionSubmitCustomExtensionHandler struct {
    OnAttributeCollectionSubmitHandler
}
// NewOnAttributeCollectionSubmitCustomExtensionHandler instantiates a new OnAttributeCollectionSubmitCustomExtensionHandler and sets the default values.
func NewOnAttributeCollectionSubmitCustomExtensionHandler()(*OnAttributeCollectionSubmitCustomExtensionHandler) {
    m := &OnAttributeCollectionSubmitCustomExtensionHandler{
        OnAttributeCollectionSubmitHandler: *NewOnAttributeCollectionSubmitHandler(),
    }
    odataTypeValue := "#microsoft.graph.onAttributeCollectionSubmitCustomExtensionHandler"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnAttributeCollectionSubmitCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnAttributeCollectionSubmitCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnAttributeCollectionSubmitCustomExtensionHandler(), nil
}
// GetConfiguration gets the configuration property value. Configuration regarding properties of the custom extension that can be overwritten per event listener.
// returns a CustomExtensionOverwriteConfigurationable when successful
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) GetConfiguration()(CustomExtensionOverwriteConfigurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomExtensionOverwriteConfigurationable)
    }
    return nil
}
// GetCustomExtension gets the customExtension property value. The customExtension property
// returns a OnAttributeCollectionSubmitCustomExtensionable when successful
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) GetCustomExtension()(OnAttributeCollectionSubmitCustomExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnAttributeCollectionSubmitCustomExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnAttributeCollectionSubmitHandler.GetFieldDeserializers()
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionOverwriteConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(CustomExtensionOverwriteConfigurationable))
        }
        return nil
    }
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnAttributeCollectionSubmitCustomExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(OnAttributeCollectionSubmitCustomExtensionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnAttributeCollectionSubmitHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. Configuration regarding properties of the custom extension that can be overwritten per event listener.
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) SetConfiguration(value CustomExtensionOverwriteConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *OnAttributeCollectionSubmitCustomExtensionHandler) SetCustomExtension(value OnAttributeCollectionSubmitCustomExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
type OnAttributeCollectionSubmitCustomExtensionHandlerable interface {
    OnAttributeCollectionSubmitHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(CustomExtensionOverwriteConfigurationable)
    GetCustomExtension()(OnAttributeCollectionSubmitCustomExtensionable)
    SetConfiguration(value CustomExtensionOverwriteConfigurationable)()
    SetCustomExtension(value OnAttributeCollectionSubmitCustomExtensionable)()
}
