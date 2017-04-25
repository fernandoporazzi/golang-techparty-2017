import React, { Component } from 'react';
import Group from './Group';

class Form extends Component {
  handleSubmit(e) {
    e.preventDefault();

    let movieForm = this.refs.createMovieForm;

    let formData = new FormData(movieForm);

    this.props.onSubmitNewMovie(formData);

    movieForm.reset();
  }

  render() {
    return (
      <form className="form" method="POST"  encType="multipart/form-data" id="createMovieForm" ref="createMovieForm" onSubmit={this.handleSubmit.bind(this)}>
        <Group type="text" name="name" label="Movie Name" id="name" />
        <Group type="number" name="year" label="Year" id="year"/>
        <Group type="file" name="cover" label="Cover Image" id="cover"/>
        <Group type="checkbox" name="category" label="Action" value="Action" id="category_action" />
        <Group type="checkbox" name="category" label="Drama" value="Drama" id="category_drama" />
        <Group type="checkbox" name="category" label="Adventure" value="Adventure" id="category_adventure" />
        <Group type="checkbox" name="category" label="Fiction" value="Fiction" id="category_fiction" />
        <Group type="checkbox" name="category" label="Thriller" value="Thriller" id="category_thriller" />
        <Group type="textarea" name="synopsis" label="Synopsis" value="Synopsis" id="synopsis"/>

        <input type="submit"/>
      </form>
    );
  }
}

export default Form;
