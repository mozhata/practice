var TestClickComponet = React.createClass({
	render: function() {
		return (
			<div>
				<button onClick={this.clickHandler}>显示|隐藏</button><span ref="tip">测试点击</span>
			</div>
		);
	},
	clickHandler: function(event) {
		var tipE = React.findDOMNode(this.refs.tip);
		if (tipE.style.display === "none"){
			tipE.style.display = "inline";
		}else{
			tipE.style.display = "none";
		}

		event.stopPropagation();
		event.preventDefault();
	},
});

var TestInputComponent = React.createClass({
	render: function() {
		return (
			<div>
				<input type="text" onChange={this.changeHandler}/><span>{this.state.inputContent}</span>
			</div>
		);
	},
	getInitialState: function() {
		return (
			{
				inputContent: ""
			}
		);
	},
	changeHandler: function(event){
		this.setState({
			inputContent: event.target.value
		});

		event.preventDefault();
		event.stopPropagation();
	},
});

React.render(
	<div>
		<TestClickComponet/>
		<br/>
		<br/>
		<TestInputComponent/>
	</div>,
	document.getElementById('container'));

var HeeloMessage = React.createClass({
	render: function(){
		return <div>hello {this.props.name}</div>
	}
});
React.render(<HeeloMessage name="John"/>, document.getElementById("hello"));

var Timer = React.createClass({
	getInitialState: function(){
		return {secondsElapsed:0};
	},
	tick:function(){
		this.setState({secondsElapsed: this.state.secondsElapsed + 1});
	},
	componentDidMount:function(){
		this.interval = setInterval(this.tick, 1000);
	},
	componentWillUnmount: function(){
		clearInterval(this.state.interval);
	},
	render: function(){
		return (
			<div>Seconds Elapsed: {this.state.secondsElapsed}</div>
		);
	},
});
React.render(<Timer />, document.getElementById("timer"));

var TodoList = React.createClass({
	render: function(){
			var creatItem = function(itemText){
				return <li>{itemText}</li>
			};
			return <ul>{this.props.items.map(creatItem)}</ul>;
		}
});
var TodoApp = React.createClass({
	render: function(){
		return(
			<div>
				<h3>TODO</h3>
				<TodoList items={this.state.items} />
				<form onSubmit={this.handleSubmit}>
					<input onChange={this.onChange} value={this.state.text} />
					<button>{'Add #' + (this.state.items.length + 1)}</button>
				</form>
			</div>
		);
	},
	getInitialState: function(){
		return {items: [], text:""};
	},
	onChange: function(e){
		this.setState({text: e.target.value});
	},
	handleSubmit: function(e){
		e.preventDefault();
		var nextItems = this.state.items.concat([this.state.text]);
		var nextText = "";
		this.setState({items: nextItems, text:nextText});
	},
});
React.render(<TodoApp/>, document.getElementById("todoapp"));


var CommentBox = React.createClass({
	render: function(){
		return (
			<div className="commentBox">
				<h1>Comments</h1>
				<CommentList />
				<CommentForm />
			</div>
		);
	}
});
React.render(
	<CommentBox />,
	document.getElementById("commentBox")
);
var CommentList = React.createClass({
	render: function(){
		return (
			<div className="commentList">
				<Comment author="Peter Hunt">this is one comment </Comment>
				<Comment author="Jordan Walke">this is *another* comment </Comment>
			</div>
		);
	},
});var CommentForm = React.createClass({
	render: function(){
		return (
			<div className="commentForm">
				Hello, world ! I am a CommentForm .
			</div>
		);
	},
});
var Comment = React.createClass({
	render: function(){
		<div className="comment">
			<h2>
				{this.props.author}
			</h2>
			{marked(this.props.children.toString())}
		</div>
	}
});