// script.js
document.getElementById("masterbookForm").addEventListener("submit", function(event) {
    event.preventDefault();
    var formData = new FormData(this);
    
    // Call the appropriate backend API endpoints
    analyzeAndDiagnose(formData);
});

function analyzeAndDiagnose(formData) {
    // Call the Python backend API
    fetch("/python/analyze", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        // Process the Python analysis results
        var pythonResult = data.result;
        // Update the UI or perform further actions based on the Python analysis result
    })
    .catch(error => {
        console.error("Error analyzing in Python:", error);
    });

    // Call the Rust backend API
    fetch("/rust/analyze", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        // Process the Rust analysis results
        var rustResult = data.result;
        // Update the UI or perform further actions based on the Rust analysis result
    })
    .catch(error => {
        console.error("Error analyzing in Rust:", error);
    });

    // Call the Go backend API
    fetch("/go/analyze", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        // Process the Go analysis results
        var goResult = data.result;
        // Update the UI or perform further actions based on the Go analysis result
    })
    .catch(error => {
        console.error("Error analyzing in Go:", error);
    });
}
