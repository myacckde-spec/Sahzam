# Sahzam

Sahzam is a compact, educational MVP for a Shazam-like audio fingerprinting workflow in Go.

It demonstrates the core idea behind music identification in a simple, readable way:

1. Load a WAV file.
2. Decode PCM samples.
3. Split the signal into windows.
4. Apply FFT.
5. Detect dominant frequencies.
6. Build a simple fingerprint.
7. Compare it with a local demo database.
8. Report the best match with similarity and confidence.

## Краткая инструкция по работе

### Как запустить

1. Установите Go 1.22+.
2. Перейдите в папку проекта:
   ```bash
   cd /workspaces/Sahzam
   ```
3. Запустите тесты:
   ```bash
   go test ./...
   ```
4. Запустите приложение на примере файла [sample.wav](sample.wav):
   ```bash
   go run ./cmd/shazam ./sample.wav
   ```

### Что делает приложение

- Считывает WAV-файл.
- Превращает аудио в набор PCM-выборок.
- Разбивает сигнал на небольшие окна.
- Применяет FFT, чтобы увидеть основные частоты.
- Формирует простой аудио-отпечаток.
- Сравнивает его с локальной базой примеров.
- Выводит наиболее похожую песню, оценку сходства и уверенность.

### Где что находится в коде

- [cmd/shazam/main.go](cmd/shazam/main.go) — точка входа, запуск всего процесса.
- [reader/wav.go](reader/wav.go) — чтение и разбор WAV.
- [audio/audio.go](audio/audio.go) — общая структура данных аудио.
- [fft/fft.go](fft/fft.go) — простая реализация FFT.
- [fingerprint/fingerprint.go](fingerprint/fingerprint.go) — модель отпечатка.
- [matcher/matcher.go](matcher/matcher.go) — сравнение отпечатков.
- [database/database.go](database/database.go) — маленькая демонстрационная база данных.
- [utils/logger.go](utils/logger.go) — логирование.

### Как это устроено по шагам

1. В [cmd/shazam/main.go](cmd/shazam/main.go) запускается главный сценарий приложения.
   - программа проверяет аргументы командной строки;
   - находит путь к WAV-файлу;
   - создает объекты для чтения, обработки и сравнения;
   - запускает цепочку: загрузка → обработка → сравнение.

2. [reader/wav.go](reader/wav.go) отвечает за чтение аудио-файла.
   - проверяет, что файл действительно имеет расширение .wav;
   - читает заголовок RIFF/WAVE;
   - получает частоту дискретизации, число каналов и глубину битности;
   - читает PCM-данные и превращает их в массив чисел.

3. [audio/audio.go](audio/audio.go) хранит модель аудио.
   - здесь описан тип AudioData, который содержит массив сэмплов, частоту дискретизации, глубину и количество каналов;
   - это общая структура, которую передают между пакетами.

4. [fft/fft.go](fft/fft.go) выполняет простое преобразование Фурье.
   - сигнал из временной области переводится в частотную;
   - это помогает понять, какие частоты доминируют в окне аудио;
   - в этом учебном проекте FFT сделана максимально просто и понятно.

5. [fingerprint/fingerprint.go](fingerprint/fingerprint.go) описывает аудио-отпечаток.
   - здесь хранится структура Fingerprint;
   - в ней фиксируются самые сильные частоты и размер окна;
   - это упрощённый аналог отпечатка, который реально используют в системах распознавания музыки.

6. [cmd/shazam/main.go](cmd/shazam/main.go) уже после чтения аудио формирует fingerprint.
   - программа разбивает сигнал на окна;
   - для каждого окна applies FFT;
   - выбирает наиболее сильную частоту;
   - сохраняет её в список peak-частот.

7. [matcher/matcher.go](matcher/matcher.go) сравнивает отпечаток с базой.
   - для каждого кандидата из базы считается, насколько сильно совпадают частоты;
   - выбирается лучший результат;
   - на основе сходства определяется уровень уверенности: High / Medium / Low.

8. [database/database.go](database/database.go) содержит демо-базу.
   - здесь лежат 3 примера песен с заранее заданными отпечатками;
   - это не настоящая база данных, а учебный набор данных для демонстрации.

9. [utils/logger.go](utils/logger.go) отвечает за вывод сообщений в консоль.
   - программа сообщает пользователю, на каком этапе она сейчас находится.

### Понять поток данных очень просто

Если смотреть по цепочке, то получается так:

```text
WAV файл
  → reader/wav.go
  → audio.AudioData
  → cmd/shazam/main.go
  → FFT
  → fingerprint.Fingerprint
  → matcher.SimpleMatcher
  → результат: песня + similarity + confidence
```

### Что именно происходит в самом коде

- В [cmd/shazam/main.go](cmd/shazam/main.go) есть структура App.
  - она объединяет все зависимости: читатель аудио, генератор fingerprint, matcher, база данных и логгер;
  - это делает код более читаемым и удобно расширяемым.

- В функции Run() происходит основной сценарий.
  - сначала логируется "Loading audio...";
  - затем читается файл;
  - после этого происходит "Decoding WAV...";
  - затем генерируется отпечаток;
  - затем выполняется поиск в базе;
  - в конце выводится лучший вариант.

- В функции Generate() в [cmd/shazam/main.go](cmd/shazam/main.go) используется простая логика.
  - берётся окно сигнала;
  - нормализуется;
  - применяется FFT;
  - определяется частота с наибольшей амплитудой;
  - она добавляется в список отпечатка.

### Почему это называется учебным проектом

Потому что он специально упрощён:

- использует маленькую локальную базу, а не огромную коллекцию песен;
- делает очень простой FFT, а не профессиональный DSP-пайплайн;
- сравнивает отпечатки по простому совпадению частот, а не по сложным хэшам и временным меткам;
- полезен для понимания идеи, а не для полноценного промышленного распознавания музыки.

## Для новичка: что запускать и как смотреть на проект

### 1. Что нужно сделать в самом начале

Если ты только начинаешь разбираться, сделай так:

```bash
cd /workspaces/Sahzam
go test ./...
go run ./cmd/shazam ./sample.wav
```

Что это делает:

- `go test ./...` — запускает тесты и проверяет, что проект не сломан;
- `go run ./cmd/shazam ./sample.wav` — запускает программу на примере файла [sample.wav](sample.wav).

### 2. Что ты должен увидеть после запуска

Программа должна вывести примерно такое:

```text
Loading audio...
Decoding WAV...
Applying FFT...
Generating fingerprint...
Searching...
Song found.

Best Match:
Imagine Dragons - Believer
Similarity: 100.0%
Confidence: High
```

Это значит, что приложение успешно:

- прочитало аудио;
- обработало его;
- сформировало отпечаток;
- нашло самый похожий пример в базе.

### 3. Что тестировать

Самое простое, что можно проверить:

- что тесты проходят (`go test ./...`);
- что приложение запускается на [sample.wav](sample.wav);
- что при изменении логики fingerprint или matcher результат остаётся понятным.

### 4. Из чего примерно состоит проект

Проект можно представить как цепочку из нескольких простых шагов:

```text
файл.wav
  → прочитать файл
  → получить числа из аудио
  → посмотреть главные частоты
  → сделать отпечаток
  → сравнить с примерами
  → выбрать лучший результат
```

### 5. Где смотреть код новичку

Если хочется понять проект без лишней сложности, начни с этих файлов:

- [cmd/shazam/main.go](cmd/shazam/main.go) — здесь всё собирается воедино;
- [reader/wav.go](reader/wav.go) — здесь читается WAV;
- [fft/fft.go](fft/fft.go) — здесь ищутся частоты;
- [matcher/matcher.go](matcher/matcher.go) — здесь сравниваются результаты.

### 6. Очень коротко по папкам

- [cmd/shazam](cmd/shazam) — запуск программы;
- [reader](reader) — чтение аудио;
- [fft](fft) — работа с частотами;
- [fingerprint](fingerprint) — отпечаток песни;
- [matcher](matcher) — поиск похожего варианта;
- [database](database) — примеры песен;
- [utils](utils) — вспомогательные вещи.

### 7. Если хочется понять проект быстрее

Сначала посмотри в таком порядке:

1. [cmd/shazam/main.go](cmd/shazam/main.go)
2. [reader/wav.go](reader/wav.go)
3. [fft/fft.go](fft/fft.go)
4. [matcher/matcher.go](matcher/matcher.go)

Это самый быстрый путь, чтобы понять, как всё работает.

## Architecture Diagram

```text
User / CLI
  │
  ▼
cmd/shazam
  │
  ├── reader      -> loads WAV audio
  ├── audio       -> shared audio model
  ├── fft         -> simple FFT implementation
  ├── fingerprint -> fingerprint structure and generator interface
  ├── matcher     -> compares fingerprints
  ├── database    -> demo fingerprint database
  └── utils       -> logging helpers
```

## Folder Structure

```text
cmd/shazam/          # CLI entrypoint
reader/              # WAV decoding
audio/               # audio data types
fft/                 # FFT helper code
fingerprint/         # fingerprint model and interfaces
database/            # demo fingerprint database
matcher/             # similarity matching logic
utils/               # logging helpers
```

## Package Overview

- cmd/shazam
  - Entry point for running the app from the command line.
- reader
  - Reads and decodes basic WAV files into PCM samples.
- audio
  - Defines the shared audio data structure.
- fft
  - Implements a small educational FFT and magnitude extraction.
- fingerprint
  - Stores fingerprints and defines the generator interface.
- matcher
  - Compares query fingerprints with database fingerprints.
- database
  - Contains a tiny local demo database of sample songs.
- utils
  - Provides simple logging support.

## Installation

Make sure Go 1.22+ is installed.

```bash
go version
```

## Build and Run

```bash
go test ./...
go run ./cmd/shazam ./sample.wav
```

## Example Output

```text
Loading audio...
Decoding WAV...
Applying FFT...
Generating fingerprint...
Searching...
Song found.

Best Match:
Imagine Dragons - Believer
Similarity: 100.0%
Confidence: High
```

## Teaching Notes

### How WAV is decoded

The WAV reader reads the RIFF/WAVE header, extracts sample rate and bit depth, and decodes the PCM payload into a slice of float64 values.

### What FFT does

FFT converts a time-domain signal into frequency-domain information. In this MVP, it helps identify the dominant frequencies present in a small audio window.

### How fingerprints work

A fingerprint is a compact summary of the signal. This MVP uses the strongest frequency peaks from several windows to build a simple signature.

### How matching works

The application compares the query fingerprint against the local demo database and returns the song whose fingerprint has the highest overlap with the query.

### How this differs from real Shazam

Real Shazam uses robust fingerprints, time offsets, hashing, large-scale indexing, and much more sophisticated DSP. This project intentionally keeps the workflow small and understandable.

## Future Improvements

This MVP can evolve into a production-grade system by adding:

- MP3 support
- spectrogram-based fingerprints
- robust hash-based matching
- PostgreSQL or Redis-backed storage
- concurrent processing with goroutines
- a REST API
- Docker packaging
- benchmarks and profiling
