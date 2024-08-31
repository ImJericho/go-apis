package model

type MainDB struct {
	AdmissionNo string `json:"admission_no,omitempty" bson:"admission_no,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	HellPoints  int    `json:"hell_points,omitempty" bson:"hell_points,omitempty"`
	HevenPoints int    `json:"heven_points,omitempty" bson:"heven_points,omitempty"`
}
