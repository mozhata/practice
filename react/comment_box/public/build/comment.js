var CommentBox = React.createClass({
	displayName: "CommentBox",

	render: function () {
		return React.createElement(
			"div",
			{ className: "commentBox" },
			"hello  world ! I am a CommentBox."
		);
	}
});

var CommentList = React.createClass({
	displayName: "CommentList",

	render: function () {
		return React.createElement(
			"div",
			{ className: "commentList" },
			"hello, world ! I am a CommentList"
		);
	}
});

ReactDOM.render(React.createElement(CommentBox, null), document.getElementById("content"));