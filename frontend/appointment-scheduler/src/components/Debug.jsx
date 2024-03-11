
function Debug(props) {
	return (
		<div style={{height: "100%",
			width: "100%",
			fontSize: "200%",
			display: "flex",
			justifyContent: "center",
			alignItems: "center",
			backgroundColor: "peru",
			color: "whitesmoke"}}>
			{props.text}
		</div>
	)
}

export default Debug;