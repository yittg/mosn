/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package java_exception

// InvocationTargetException represents an exception of the same name in java
type InvocationTargetException struct {
	SerialVersionUID     int64
	DetailMessage        string
	SuppressedExceptions []Throwabler
	StackTrace           []StackTraceElement
	Cause                Throwabler
	Target               Throwabler
}

// NewInvocationTargetException is the constructor
func NewInvocationTargetException(target Throwabler, detailMessage string) *InvocationTargetException {
	return &InvocationTargetException{DetailMessage: detailMessage, StackTrace: []StackTraceElement{}, Target: target}
}

// Error output error message
func (e InvocationTargetException) Error() string {
	return e.DetailMessage
}

// JavaClassName  java fully qualified path
func (InvocationTargetException) JavaClassName() string {
	return "java.lang.reflect.InvocationTargetException"
}
