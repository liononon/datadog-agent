// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package providers

import (
	"time"

	"github.com/DataDog/datadog-agent/pkg/config"
)

// syncTimeout can be used to wait for the kubernetes client-go cache to sync.
var syncTimeout = config.Datadog.GetDuration("cache_sync_timeout") * time.Second
