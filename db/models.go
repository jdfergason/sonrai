/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package db

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Jobs []*Job
}

type JobResponse struct {
	TraceID string
	Code int
	Status string
	Data []*Job
}

type Job struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	Kind        string
	Command     string
	Arguments   []string
	Environment map[string]string
	Schedule    string
	LastRun     time.Time
	Tags        []string
}

type JobRun struct {

}

type JobLog struct {
	
}

type Secret struct {
	Name   string
	Secret string
}
