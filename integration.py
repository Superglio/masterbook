from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html')

# Add more routes and functionalities as needed

if __name__ == '__main__':
    app.run(debug=True)
