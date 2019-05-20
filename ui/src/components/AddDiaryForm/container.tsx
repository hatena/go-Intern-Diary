import React from "react";
import {MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";

import { CreateDiary } from "../__generated__/CreateDiary"
import { GetVisitor } from "../__generated__/GetVisitor";

import { query as getVisitorQuery } from "../index"
import { CreateDiaryForm } from "./addDiary";

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
}

interface DiaryFormState {
    name: string;
    tags: Tag[];
}

export type Tag = {
    name: string
}

export class CreateDiaryFormContainer extends React.PureComponent<DiaryFormProps, DiaryFormState> {
    
    state = {
        name: "",
        tags: [],
    }

    private handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.currentTarget;
        this.setState({
            name: input.value
        })
    }

    private handleSubmit = (create: () => void) => (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        create()
        this.setState({
            name: ""
        })
    }

    render() {
        return (
            <CreateDiaryForm 
                name={this.state.name} 
                handleSubmit={this.handleSubmit} 
                handleInput={this.handleInput}
                tags={[]} 
            />
        )

    }
}