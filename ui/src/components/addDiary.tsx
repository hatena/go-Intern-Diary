import React from "react";
import {Mutation, MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";
import { diaryFragment } from "./diary";

import { CreateDiary } from "./__generated__/CreateDiary"
import { GetVisitor } from "./__generated__/GetVisitor";

import { query as getVisitorQuery } from "./index"

const createDiaryFragment = gql`
    fragment createDiaryFragment on Diary {
        id
        name
    }
`;

export const mutation = gql`
    mutation CreateDiary($name: String!) {
        createDiary(name: $name) {
            ...createDiaryFragment
        }
    }
    ${createDiaryFragment}
`

export const createUpdateDiary: MutationUpdaterFn<CreateDiary> = (cache, result) => {
    const { data } = result;
    const visitor = cache.readQuery<GetVisitor>({ query: getVisitorQuery })
    if (visitor && data) {
        const diaries = [...visitor.visitor.diaries]
        diaries.unshift(data.createDiary)
        const newVisitor = {
            visitor: {
                ...visitor.visitor,
                diaries,
            }
        }
        cache.writeQuery({query: getVisitorQuery, data: newVisitor})
    }
}

interface DiaryFormProps {
    create: (name: string) => void
}

interface DiaryFormState {
    name: string;
}
export class CreateDiaryForm extends React.PureComponent<DiaryFormProps, DiaryFormState> {
    
    state = {
        name: "",
    }

    private handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.currentTarget;
        this.setState({
            name: input.value
        })
    }

    private handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        this.props.create(this.state.name)
        this.setState({
            name: ""
        })
    }

    render() {
        return (
            <form className="CreateDiaryForm" onSubmit={this.handleSubmit}>
                <div>
                    <label>Diary Name:
                        <input type="TEXT" name="name" value={this.state.name} onChange={this.handleInput} />
                    </label>
                </div>
                <div>
                    <button>Create New</button>
                </div>
            </form>
        )

    }
}
