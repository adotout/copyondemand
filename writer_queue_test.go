/* This file is part of copyondemand.
 *
 * Copyright © 2020 Datto, Inc.
 * Author: Bryan Ehrlich <behrlich@datto.com>
 *
 * Licensed under the Apache Software License, Version 2.0
 * Fedora-License-Identifier: ASL 2.0
 * SPDX-2.0-License-Identifier: Apache-2.0
 * SPDX-3.0-License-Identifier: Apache-2.0
 *
 * copyondemand is free software.
 * For more information on the license, see LICENSE.
 * For more information on free software, see <https://www.gnu.org/philosophy/free-sw.en.html>.
 *
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package copyondemand

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueueAndDequeue(t *testing.T) {
	q := NewWriterQueue()

	wa := &QueuedWriteAction{}
	wa.actionType = 1337

	q.TryEnqueue(wa, 0)

	dq := q.TryDequeue(100, false)

	assert.Equal(t, dq.actionType, WriteActionType(1337))
}

func TestDequeueTimesOut(t *testing.T) {
	q := NewWriterQueue()

	dq := q.TryDequeue(10, false)

	assert.Nil(t, dq)
}
