package database

import "sahzam/fingerprint"

// DemoDatabase stores a tiny set of sample fingerprints.
type DemoDatabase struct {
	Records []fingerprint.FingerprintRecord
}

// NewDemoDatabase creates the sample database.
func NewDemoDatabase() *DemoDatabase {
	return &DemoDatabase{Records: []fingerprint.FingerprintRecord{
		{Song: "Imagine Dragons - Believer", Fingerprint: &fingerprint.Fingerprint{TopFrequencies: []int{440, 880, 1320, 1760}, WindowSize: 1024}},
		{Song: "The Beatles - Hey Jude", Fingerprint: &fingerprint.Fingerprint{TopFrequencies: []int{330, 660, 990, 1320}, WindowSize: 1024}},
		{Song: "Daft Punk - Around the World", Fingerprint: &fingerprint.Fingerprint{TopFrequencies: []int{220, 440, 660, 1100}, WindowSize: 1024}},
	}}
}

// Records returns the database entries.
func (d *DemoDatabase) RecordsList() []fingerprint.FingerprintRecord {
	return d.Records
}
