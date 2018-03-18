import * as React from "react";
import { Card } from "antd";
import "whatwg-fetch";

interface TodoItem {
    id: number;
    name: string;
    isCompleted: boolean;
}

interface IHomeState {
    todoItems: TodoItem[];
}

class HomePage extends React.Component<{}, IHomeState> {
    constructor(props: {}) {
        super(props);
        this.state = {todoItems: []};
        this.mapTodoItem = this.mapTodoItem.bind(this);
        this.mapTodoItems = this.mapTodoItems.bind(this);
    }

    public componentDidMount() {
        // debugger;
        fetch(`/todo`)
            .then((result) => (result.json()))
            .then(this.mapTodoItems)
            .then((todoItems) => {
                this.setState({todoItems});
            });
    }

    public render(): JSX.Element {
        return (
            <Card bordered title="Welcome to Go Todo" style={{ margin: "16px 16px"}}>
                <p>{this.state.todoItems.length > 0 ? this.state.todoItems[0].name : "welcome"}</p>
            </Card>
        );
    }

    private mapTodoItems(items: any[]): TodoItem[] {
        return items.map(this.mapTodoItem);
    }

    private mapTodoItem(item: any): TodoItem {
        return {
            id: item.id,
            name: item.name,
            isCompleted: item.isCompleted,
        };
    }
}

export default HomePage;
