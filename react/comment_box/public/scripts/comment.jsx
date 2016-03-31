var CommentBox = React.createClass({
	render: function(){
		return (
			<div className="commentBox">
				hello  world ! I am a CommentBox.
			</div>
		);
	}
});

var CommentList = React.createClass({
	render: function(){
		return (
			<div className="commentList">
				hello, world ! I am a CommentList
			</div>
		);
	}
});

ReactDOM.render(
	<CommentBox/>,
	document.getElementById("content")
);