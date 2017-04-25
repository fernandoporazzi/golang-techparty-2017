import React, { Component } from 'react';
import Sidebar from './sidebar/Sidebar';
import View from './view/View';
import './App.css';

class App extends Component {

  constructor() {
    super();

    this.state = {
      movies: []
    }
  }

  componentDidMount() {
    this.fetchData();
  }

  fetchData() {
    var self = this;

    fetch('http://localhost:8001/movie')
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        self.setState({
          movies: data
        });
      })
      .catch((err) => {
        console.log(err);
      })
  }

  handleSubmitNewMovie(formData) {
    console.log(formData);
    let self = this;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:8001/movie', true);
    xhr.onload = function(e) {
      let resp = JSON.parse(xhr.responseText);

      self.setState({movies: self.state.movies.concat(resp)});
    };

    xhr.send(formData);
  }

  render() {
    return (
      <div className="App">
        <Sidebar onSubmitNewMovie={this.handleSubmitNewMovie.bind(this)}/>
        <View movies={this.state.movies} />
      </div>
    );
  }
}

export default App;
