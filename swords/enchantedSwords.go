package swords

type EnchantedSword struct {
	// Embed the Sword type
	Sword
}

// Damage returns the damage dealt by the enchanted sword.
func (EnchantedSword) Damage() int {
	return 42
}
