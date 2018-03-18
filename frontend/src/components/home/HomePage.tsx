import * as React from "react";
import { Card, Table } from "antd";
import { ColumnProps } from "antd/lib/table";
import "whatwg-fetch";

interface TodoItem {
    id: number;
    key: number;
    name: string;
    isCompleted: boolean;
}

interface IHomeState {
    todoItems: TodoItem[];
}

const columns: Array<ColumnProps<TodoItem>> = [{
    title: "Id",
    dataIndex: "id",
    key: "id",
}, {
    title: "Name",
    dataIndex: "name",
    key: "name",
}, {
    title: "Completed",
    dataIndex: "isCompleted",
    key: "isCompleted",
    render: (text: any, record: TodoItem, index: number) => {
        return <span>{record.isCompleted ? "true" : "false"}</span>;
    },
}];

class HomePage extends React.Component<{}, IHomeState> {
    constructor(props: {}) {
        super(props);
        this.state = {todoItems: []};
        this.mapTodoItem = this.mapTodoItem.bind(this);
        this.mapTodoItems = this.mapTodoItems.bind(this);
    }

    public componentDidMount() {
        const url = `http://localhost:8080/todo`;  // for local debugging
        // const url = `/todo`;   // for deployment
        fetch(url)
            .then((result) => (result.json()))
            .then(this.mapTodoItems)
            .then((todoItems) => {
                this.setState({todoItems});
            });
    }

    public render(): JSX.Element {
        return (
            <Card bordered title="Welcome to Go Todo" style={{ margin: "16px 16px"}}>
                <Table dataSource={this.state.todoItems} columns={columns} />
            </Card>
        );
    }

    private mapTodoItems(items: any[]): TodoItem[] {
        // return items.map(this.mapTodoItem);
        return items.map(item => {
            return item as TodoItem;
        });
    }

    private mapTodoItem(item: any): TodoItem {
        // return {
        //     id: item.id,
        //     key: item.key,
        //     name: item.name,
        //     isCompleted: item.isCompleted,
        // };
        return item as TodoItem;
    }
}

export default HomePage;
