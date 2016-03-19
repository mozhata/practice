var TestClickComponet = React.createClass({
	displayName: "TestClickComponet",

	render: function () {
		return React.createElement(
			"div",
			null,
			React.createElement(
				"button",
				{ onClick: this.clickHandler },
				"显示|隐藏"
			),
			React.createElement(
				"span",
				{ ref: "tip" },
				"测试点击"
			)
		);
	},
	clickHandler: function (event) {
		var tipE = React.findDOMNode(this.refs.tip);
		if (tipE.style.display === "none") {
			tipE.style.display = "inline";
		} else {
			tipE.style.display = "none";
		}

		event.stopPropagation();
		event.preventDefault();
	}
});

var TestInputComponent = React.createClass({
	displayName: "TestInputComponent",

	render: function () {
		return React.createElement(
			"div",
			null,
			React.createElement("input", { type: "text", onChange: this.changeHandler }),
			React.createElement(
				"span",
				null,
				this.state.inputContent
			)
		);
	},
	getInitialState: function () {
		return {
			inputContent: ""
		};
	},
	changeHandler: function (event) {
		this.setState({
			inputContent: event.target.value
		});

		event.preventDefault();
		event.stopPropagation();
	}
});

React.render(React.createElement(
	"div",
	null,
	React.createElement(TestClickComponet, null),
	React.createElement("br", null),
	React.createElement("br", null),
	React.createElement(TestInputComponent, null)
), document.getElementById('container'));

var HeeloMessage = React.createClass({
	displayName: "HeeloMessage",

	render: function () {
		return React.createElement(
			"div",
			null,
			"hello ",
			this.props.name
		);
	}
});
React.render(React.createElement(HeeloMessage, { name: "John" }), document.getElementById("hello"));

var Timer = React.createClass({
	displayName: "Timer",

	getInitialState: function () {
		return { secondsElapsed: 0 };
	},
	tick: function () {
		this.setState({ secondsElapsed: this.state.secondsElapsed + 1 });
	},
	componentDidMount: function () {
		this.interval = setInterval(this.tick, 1000);
	},
	componentWillUnmount: function () {
		clearInterval(this.state.interval);
	},
	render: function () {
		return React.createElement(
			"div",
			null,
			"Seconds Elapsed: ",
			this.state.secondsElapsed
		);
	}
});
React.render(React.createElement(Timer, null), document.getElementById("timer"));

var TodoList = React.createClass({
	displayName: "TodoList",

	render: function () {
		var creatItem = function (itemText) {
			return React.createElement(
				"li",
				null,
				itemText
			);
		};
		return React.createElement(
			"ul",
			null,
			this.props.items.map(creatItem)
		);
	}
});
var TodoApp = React.createClass({
	displayName: "TodoApp",

	render: function () {
		return React.createElement(
			"div",
			null,
			React.createElement(
				"h3",
				null,
				"TODO"
			),
			React.createElement(TodoList, { items: this.state.items }),
			React.createElement(
				"form",
				{ onSubmit: this.handleSubmit },
				React.createElement("input", { onChange: this.onChange, value: this.state.text }),
				React.createElement(
					"button",
					null,
					'Add #' + (this.state.items.length + 1)
				)
			)
		);
	},
	getInitialState: function () {
		return { items: [], text: "" };
	},
	onChange: function (e) {
		this.setState({ text: e.target.value });
	},
	handleSubmit: function (e) {
		e.preventDefault();
		var nextItems = this.state.items.concat([this.state.text]);
		var nextText = "";
		this.setState({ items: nextItems, text: nextText });
	}
});
React.render(React.createElement(TodoApp, null), document.getElementById("todoapp"));

var CommentBox = React.createClass({
	displayName: "CommentBox",

	render: function () {
		return React.createElement(
			"div",
			{ className: "commentBox" },
			React.createElement(
				"h1",
				null,
				"Comments"
			),
			React.createElement(CommentList, null),
			React.createElement(CommentForm, null)
		);
	}
});
React.render(React.createElement(CommentBox, null), document.getElementById("commentBox"));
var CommentList = React.createClass({
	displayName: "CommentList",

	render: function () {
		return React.createElement(
			"div",
			{ className: "commentList" },
			React.createElement(
				Comment,
				{ author: "Peter Hunt" },
				"this is one comment "
			),
			React.createElement(
				Comment,
				{ author: "Jordan Walke" },
				"this is *another* comment "
			)
		);
	}
});var CommentForm = React.createClass({
	displayName: "CommentForm",

	render: function () {
		return React.createElement(
			"div",
			{ className: "commentForm" },
			"Hello, world ! I am a CommentForm ."
		);
	}
});
var Comment = React.createClass({
	displayName: "Comment",

	render: function () {
		React.createElement(
			"div",
			{ className: "comment" },
			React.createElement(
				"h2",
				null,
				this.props.author
			),
			marked(this.props.children.toString())
		);
	}
});