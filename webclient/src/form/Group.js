import React, { Component } from 'react';

class Group extends Component {
  render() {
    const l = <label htmlFor={this.props.id}>{this.props.label}</label>;

    let i;

    if (this.props.type === 'textarea') {
      i = <textarea name={this.props.name} />
    } else {
      i = <input type={this.props.type} name={this.props.name} id={this.props.id} value={this.props.value} />
    }

    return (
      <div className={`form-group ${this.props.type}`}>
        {this.props.type !== 'checkbox' ? l : ''}

        {i}

        {this.props.type === 'checkbox' ? l : ''}
      </div>
    );
  }
}

export default Group;
