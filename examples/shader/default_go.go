// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package main

var default_go = []byte("// Copyright 2020 The Ebiten Authors\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//     http://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n// +build ignore\n\npackage main\n\nvar Time float\nvar Cursor vec2\nvar ImageSize vec2\n\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\n\tpos := position.xy/textureDstSize() + Cursor/textureDstSize()/4\n\tclr := 0.0\n\tclr += sin(pos.x*cos(Time/15)*80) + cos(pos.y*cos(Time/15)*10)\n\tclr += sin(pos.y*sin(Time/10)*40) + cos(pos.x*sin(Time/25)*40)\n\tclr += sin(pos.x*sin(Time/5)*10) + sin(pos.y*sin(Time/35)*80)\n\tclr *= sin(Time/10) * 0.5\n\treturn vec4(clr, clr*0.5, sin(clr+Time/3)*0.75, 1)\n}\n")
