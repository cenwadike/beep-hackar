package types

// ValidateBasic is used for validating the packet
func (p IntentPacketPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IntentPacketPacketData) GetBytes() ([]byte, error) {
	var modulePacket IntentPacketData

	modulePacket.Packet = &IntentPacketData_IntentPacketPacket{&p}

	return modulePacket.Marshal()
}
