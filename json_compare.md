# Json Lib Performance Comparison

| lib | Small(Speed) | Small(Memory) | Medium(S) | Medium(M) | Large(S) | Large(M) | Large payload(S) | Large payload(M) | Medium encode(S) | Medium encode(M) | Easy of use |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| jsoniter | ✅ | ✅ | ⚠️ | ⚠️ | ✅ | ✅ | ⚠️ | ⛔ | ✅ | ✅ | OK |
| easyjson | ✅ | ✅ | ⛔ | ✅ | ⚠️ | ⚠️ | --- | --- | ✅ | ✅ | HARD |
| go-json | ✅ | ⚠️ | ⚠️ | ⛔ | ⚠️ | ⛔ | ✅ | ⛔ | ✅ | ✅ | **EASY** |
| jsonparser | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | --- | --- | HARD |
| fastjson | ⛔️ | ⛔ | ⛔ | ⛔ | --- | --- | --- | --- | --- | --- | **EASY** |
| simdjson | ⛔️ | ⛔ | --- | --- | --- | --- | --- | --- | --- | --- | HARD |
| encode/json | ⛔️  | ⚠️ | ⛔️  | ⚠️ | ⛔️ | ⚠️ | ⛔️ | ✅ | ✅ | ✅ | **EASY** |

✅ Good Performance, very small margin between all green marks   
⚠️ Ok Performance, not as good as the green marks   
⛔️ Worst   