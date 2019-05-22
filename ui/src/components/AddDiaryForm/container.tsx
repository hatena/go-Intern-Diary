import React from "react";
import {MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";

import { CreateDiary } from "./__generated__/CreateDiary"
import { GetVisitor } from "../UserTop/__generated__/GetVisitor";

import { query as getVisitorQuery } from "../UserTop/container"
import { CreateDiaryForm } from "./addDiary";
import { triggerAsyncId } from "async_hooks";
import { AddTagForm } from "./addTagForm";

const createDiaryFragment = gql`
    fragment createDiaryFragment on Diary {
        id
        name
    }
`;

export const mutation = gql`
    mutation CreateDiary($name: String!, $tags: [String!]!) {
        createDiary(name: $name, tags: $tags) {
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
    tagName: string;
    tags: Tag[];
}

export type Tag = {
    name: string
}

export class CreateDiaryFormContainer extends React.PureComponent<DiaryFormProps, DiaryFormState> {
    
    constructor(props: DiaryFormProps) {
        super(props)
        const tags: Tag[] = []
        this.state = {
            name: "",
            tagName: "",
            tags: tags,
        }

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
            name: "",
            tagName: "",
        })
    }


    private handleTagInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.currentTarget;
        this.setState({
            tagName: input.value
        })
    }

    private handleTagSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        const newTag: Tag = { name: this.state.tagName }
        const updatedTags = [newTag].concat(this.state.tags)
        this.setState({
            tags: updatedTags,
            tagName: "",
        })
    }

    private handleDeleteButton = (selectedTag: Tag) => () => {
        const newTags = this.state.tags.filter(tag => tag.name != selectedTag.name)
        this.setState({
            tags: newTags
        })
    }

    render() {
        return (
            <div>
                <CreateDiaryForm 
                    name={this.state.name} 
                    handleSubmit={this.handleSubmit} 
                    handleInput={this.handleInput}
                    tags={this.state.tags}
                />
                <AddTagForm 
                    tagName={this.state.tagName}
                    tags={this.state.tags}
                    handleDeleteButton={this.handleDeleteButton}
                    handleTagInput={this.handleTagInput}
                    handleTagSubmit={this.handleTagSubmit}
                />
            </div>
        )

    }
}