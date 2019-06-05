import React from "react";
import {MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";

import { CreateDiary } from "./__generated__/CreateDiary"
import { GetVisitor } from "../UserTop/__generated__/GetVisitor";

import { query as getVisitorQuery } from "../UserTop/container"
import { CreateDiaryForm } from "./addDiary";
import { AddTagForm } from "./addTagForm";

const createDiaryFragment = gql`
    fragment createDiaryFragment on Diary {
        id
        name
    }
`;

const listCategoriesFragment = gql`
    fragment ListCategoriesFragment on Category {
        id
        category_name
    }
`

export const mutation = gql`
    mutation CreateDiary($name: String!, $tagWithCategories: [TagWithCategoryInput!]!) {
        createDiary(name: $name, tagWithCategories: $tagWithCategories) {
            ...createDiaryFragment
        }
    }
    ${createDiaryFragment}
`

export const listCategoriesQuery = gql`
    query ListCategories {
        listCategories {
            ...ListCategoriesFragment
        }
    }
    ${listCategoriesFragment}
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
    selectedCategory?: Category;
    tags: Tag[];
}

export type Tag = {
    name: string;
    category: Category;
}

export type Category = {
    id: number;
    name: string;
}

export class CreateDiaryFormContainer extends React.PureComponent<DiaryFormProps, DiaryFormState> {
    
    constructor(props: DiaryFormProps) {
        super(props)
        const tags: Tag[] = []
        const dummyCategoy: Category = {id: 0, name: "dummy"}
        this.state = {
            name: "",
            tagName: "",
            selectedCategory: dummyCategoy,
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
        
        if (this.state.name == "") {
            alert("日記の名前を入力してください")
            return 
        }

        create()
        const tags: Tag[] = []
        this.setState({
            // selectedCategory: undefined,
            name: "",
            tagName: "",
            tags: tags
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

        if (this.state.selectedCategory == undefined) {
            alert("カテゴリーが選択されていません")
            return 
        }

        if (this.state.tagName== "") {
            alert("タグが入力されていません")
            return 
        }

        const newTag: Tag = { name: this.state.tagName , category: this.state.selectedCategory}
        const updatedTags = [newTag].concat(this.state.tags)
        this.setState({
            // selectedCategory: undefined,
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

    private handleCategorySelectButton = (categories: Category[], selectedCategoryId: number) => () => {
        const selectedCategory = categories.filter(category => category.id == selectedCategoryId)[0]
        this.setState({
            selectedCategory: selectedCategory
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
                    handleCategorySelectButton={this.handleCategorySelectButton}
                    selectedCategory={this.state.selectedCategory}
                />
            </div>
        )

    }
}