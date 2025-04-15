# Phishing Email Classifier

A web application written in Go that uses machine learning to detect phishing emails. This project provides an easy-to-use interface for analyzing email content and determining whether it's a legitimate message or a phishing attempt.

## Features

- **Pure Go Implementation**: Built entirely with Go, including the machine learning components
- **Machine Learning Classification**: Uses Random Forest algorithm to identify phishing patterns
- **Text Processing**: Handles email-specific text features like URLs and email addresses
- **Web Interface**: Simple and intuitive UI for submitting and analyzing emails
- **Real-time Analysis**: Instant classification results with confidence scores
- **Key Indicators**: Shows important features that influenced the classification decision

## Screenshot

![Phishing Classifier Screenshot](https://example.com/screenshot.png)

## Requirements

- Go 1.16 or higher
- GoLearn package and its dependencies

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/phishing-classifier.git
   cd phishing-classifier
   ```

2. Initialize the Go module:
   ```bash
   go mod init phishing-classifier
   ```

3. Install the required dependencies:
   ```bash
   go get github.com/sjwhitworth/golearn/...
   ```

4. Create the necessary directories:
   ```bash
   mkdir -p data static templates
   ```

5. Download the dataset:
   - Download the Enron Phishing Corpus from [Kaggle](https://www.kaggle.com/datasets/rtatman/fraudulent-email-corpus)
   - Format it as a CSV with at least two columns (email text and label)
   - Save it to `data/phishing_dataset.csv`

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## How It Works

1. **Text Preprocessing**: Emails are cleaned and normalized (removing URLs, converting to lowercase, etc.)
2. **Feature Extraction**: The system identifies important words and patterns in the email
3. **Classification**: The Random Forest classifier analyzes the features to determine if the email is phishing
4. **Results Display**: The system shows the classification result with a confidence score and key indicators

## Project Structure

```
phishing-classifier/
├── main.go                  # Main application code
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksums
├── data/
│   └── phishing_dataset.csv # Training dataset
├── static/
│   └── styles.css           # CSS styling for the web interface
└── templates/
    ├── index.html           # Main page template
    └── result.html          # Results page template
```

## Technical Details

### Machine Learning

The system uses the GoLearn library to implement:
- Text vectorization (converting text to numerical features)
- Random Forest classification (ensemble of decision trees)
- Feature importance analysis

### Text Processing

Email-specific text processing includes:
- Replacing email addresses with tokens
- Replacing URLs with tokens
- Removing stop words (common words like "the", "and", etc.)
- Text normalization and tokenization

### Web Interface

Built using Go's standard library:
- `net/http` for the web server
- HTML templates for the UI
- No external web frameworks required

## Extending the Project

Here are some ways you could extend this project:

1. **Header Analysis**: Add analysis of email headers for additional signals
2. **URL Inspection**: Deep analysis of URLs contained in emails
3. **Model Persistence**: Save trained models to disk for faster startup
4. **API Endpoints**: Add REST API for programmatic access
5. **Improved Visualization**: Enhanced UI with charts showing feature importance
6. **Multiple Models**: Implement multiple classification algorithms for comparison

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- The Enron Phishing Corpus dataset
- GoLearn machine learning library for Go
- The Go community for their excellent packages and documentation
