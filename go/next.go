/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package runtime

import (
	"context"
	"github.com/tetratelabs/wazero/api"
)

func (r *Runtime[T]) next(ctx context.Context, module api.Module, params []uint64) {
	pointer := uint32(params[0])
	length := uint32(params[1])
	r.modulesMu.RLock()
	m := r.modules[module.Name()]
	r.modulesMu.RUnlock()

	if m == nil {
		return
	}

	buf, ok := m.module.Memory().Read(pointer, length)
	if !ok {
		return
	}

	err := m.instance.runtimeContext().Read(buf)
	if err != nil {
		return
	}

	var ctxBuffer []byte
	if m.function.next == nil {
		m.instance.ctx, err = m.instance.next(m.instance.Context())
		if err != nil {
			ctxBuffer = m.instance.runtimeContext().Error(err)
		} else {
			ctxBuffer = m.instance.runtimeContext().Write()
		}
	} else {
		err = m.function.next.Run(ctx, m.instance)
		if err != nil {
			ctxBuffer = m.instance.runtimeContext().Error(err)
		} else {
			ctxBuffer = m.instance.runtimeContext().Write()
		}
	}

	ctxBufferLength := uint64(len(ctxBuffer))
	writeBuffer, err := m.resize.Call(ctx, ctxBufferLength)
	if err != nil {
		return
	}
	module.Memory().Write(uint32(writeBuffer[0]), ctxBuffer)

	return
}
