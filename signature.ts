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

export type NewSignature<T extends Signature> = () => T;

export interface Signature {
  RuntimeContext(): RuntimeContext
}

export interface Context {
  GuestContext(): GuestContext
}

export interface RuntimeContext {
  Read(b: Uint8Array): Error | undefined;
  Write(): Uint8Array;
  Error(err: Error): Uint8Array;
}

export interface GuestContext {
  FromReadBuffer(): Error | undefined;
  ToWriteBuffer(): number[];
  ErrorWriteBuffer(err: Error): number[];
}