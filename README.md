# go-word-generator

This is an application that generates words based on pre-defined rules.

## Run this application

```bash
go build -o ./build/
./build/go-word-generator.exe [-f <path to config JSON>]

2025/02/22 22:47:38 Successfully loaded the config!
2025/02/22 22:47:38 7 sounds(s) are defined in the config.
2025/02/22 22:47:38 4 pattern(s) are defined in the config.
2025/02/22 22:47:38 Generating pattern: C_START/V/C_INTER/$
2025/02/22 22:47:38 Generating pattern: C1/V/C1/$
2025/02/22 22:47:38 Generating pattern: V/CC/$
2025/02/22 22:47:38 Generating pattern: C_NO_DIACRITICS/V/C_NO_DIACRITICS/$
gnuvra
luwan
istan
pkofa
```

## Config structure

### Sounds

Defines the sound collections used in `patterns`.
Sounds are dictionary of sound name and its sound collections.

For example, the example below defines `C` sound collection, and this collection have `k` and `g`.

```json5
{
  "C": [
    "k", "g"
  ]
}
```


### Patterns

Defines the patterns of word to be generated.
Patterns are defined as array, and each element should have `label` and `syllable` field.

* `label`: Description of this pattern.
* `syllable`: Array of sounds.

For example,

```json5
[
  {
    "label": "C/C",
    "syllable": ["C", "C"]
  }
]
```

### Example

This config example will generate words like 'kaga', 'gato', 'didi' or something else.

```json5
{
  "sounds": {
    "C": [
      // Consonants
      "k", "g", "t", "d"
    ],
    "V": [
      // Vowels
      "a", "e", "i", "o", "u"
    ],
  },
  "patterns": [
    {
      "label": "C/V/C/V words",
      "syllable": [
        "C", "V", "C", "V"
      ]
    }
  ]
}
```

## Build and run test

```bash
# Build this application
go build -o ./build/

# Run tests
go test
```

## License

Apache-2.0