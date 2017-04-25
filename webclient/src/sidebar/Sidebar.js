import React, { Component } from 'react';
import Form from '../form/Form';

class Sidebar extends Component {
  render() {
    return (
      <aside className="sidebar">
        <h1>Movie API</h1>
        <h4>Powered by GoLang</h4>

        <Form onSubmitNewMovie={this.props.onSubmitNewMovie}/>
      </aside>
    );
  }
}

export default Sidebar;
