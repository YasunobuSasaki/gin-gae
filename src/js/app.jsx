import React from 'react';
import ReactDOM from 'react-dom';

var createReactClass = require('create-react-class');
//var React = require('react');

// コンポーネント名の頭文字は大文字にする
var HelloReact = createReactClass({
    render: function () {
        return (
            <div>
                Hello World!!!
            </div>
        );
    }
});

ReactDOM.render(
    <HelloReact />,
    document.getElementById('content')
);