package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Enpoints struct {
	ID            primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceName   string              `json:"service_name,omitempty" bson:"service_name,omitempty"`
	ServicePrefix string              `json:"service_prefix,omitempty" bson:"service_prefix,omitempty"`
	ServiceHost   string              `json:"service_host,omitempty" bson:"service_host,omitempty"`
	ServicePort   string              `json:"service_port,omitempty" bson:"service_port,omitempty"`
	CreatedAt     *primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt     *primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}

type EnpointsRepository interface {
	FindByServicePrefix(servicePrefix string) (Enpoints, error)
}
