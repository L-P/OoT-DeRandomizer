package main

// QuestTypeMaster makes all dungeons use the MQ variant.
const QuestTypeMaster string = "master"

// QuestTypeMixed randomly replaces dungeons with their MQ variant.
const QuestTypeMixed string = "mixed"

// QuestTypeNormal keeps good ol' dungeons.
const QuestTypeNormal string = "normal"

// Settings holds the ROM generation settings.
type Settings struct {
	Quest string

	// Ignore Gold Skulltullas
	IgnoreGS bool
}
