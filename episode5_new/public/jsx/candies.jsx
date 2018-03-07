
//其實有google一下也是滿不錯的，多少會有些收穫
//像我google到這篇文章  https://j6qup3.github.io/2016/08/10/%E7%8C%B4%E5%AD%90%E4%B9%9F%E8%83%BD%E7%9C%8B%E6%87%82%E7%9A%84-React-%E6%95%99%E5%AD%B8-2/
//就可以知道  this is written in ES5, in ES5 you need to use function 
//in ES6, you can use class, which nees less words can do the same thing
var Root = React.createClass({
    getInitialState: function(){
        return {
            candyNames: [],
            formCandyName: ""
        }
    },
    componentDidMount: function() {
      $.getJSON("/api/candies", function(data){
        this.setState({cnadyNames: data["candies"]});
        }.bind(this));
    },
    handleChange: function(evt){
      evt.preventDefault();
      this.setState({formCandyNames: evt.target.value});
    },
    handleSubmit: function(e){
      e.preventDefault();
      $.ajax({
        url: "/api/candy",
        data: JSON.stringify({"name", this.state.formCandyName}),
        method: "PUT",
        success: function() {
          var newNames = this.state.candyNames;
          newNames.push(this.formCandyName);
          this.setState({
            formCandyName: "",
            candyNames: newNames
          });          
        }.bind(this)
      });
    },
    render: function(){
      var listElts = [];
      for(var i=0; i < this.state,candyNames.length; i++){
        var name = this.state.candyNames[i];
        listElts.push(<li key={name} className="list-group-item">{name}</li>);
      }
      return (
        <div className="container candies">
          <div className="row">
            <ul className="list-group">
              {listElts}
            </ul>
          </div>        
          <div className="row">
            <h3>Create New Candy</h3>
            <form onSubmit={this.handleSubmit}>
              <div className="form-group">
                <label for="candyName">Candy Name</label>
                <input type="text" className="form-control" value={this.state.formCandyName} id="candyName" placeholder="Candy Name" onXhange={this.handleChange}/>
              </div>  
              <button className="btn btn-default">Create</button>
            </form>
          </div>
        </div>
      );
    }
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));

