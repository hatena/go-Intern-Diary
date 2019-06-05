import React from "react"

import {Tag, Category} from "./container"
import { Query } from "react-apollo";

import {listCategoriesQuery as query} from "./container"
import {ListCategories} from "./__generated__/ListCategories"

interface AddTagFormProps {
    tagName: string
    handleTagSubmit: (e: React.FormEvent<HTMLFormElement>) => void
    handleTagInput: (e: React.ChangeEvent<HTMLInputElement>) => void
    handleDeleteButton: (selectedTag: Tag) => () => void
    tags: Tag[]
    selectedCategory?: Category
    handleCategorySelectButton: (categories: Category[], selectedCategoryId: number) => () => void
}

export const AddTagForm: React.StatelessComponent<AddTagFormProps> = 
({
    tagName, 
    handleTagSubmit, 
    handleTagInput, 
    handleDeleteButton, 
    tags, 
    selectedCategory, 
    handleCategorySelectButton
}) => {
    return (
        <Query<ListCategories> query={query}>
            {result => {
                if (result.error) {
                    return <p className="error">Error: {result.error.message}</p>
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const {data} = result;
                const categories: Category[] = data!.listCategories.map(c => ({id: c.id, name: c.category_name}))
                return (
                    <div>
                        <div>
                        {tags.map(tag =>
                            <span key={tag.name}>
                                <button onClick={handleDeleteButton(tag)}>x: <h3>{tag.name}</h3></button>
                            </span>
                        )}
                        </div>
                        <form className="AddTagForm" onSubmit={handleTagSubmit}>
                            <div>
                                <label>Tag Name:
                                    <input type="TEXT" name="name" value={tagName} onChange={handleTagInput} />
                                </label>
                            </div>
                            {/* <div>choose category for this tag: 
                                <div>
                                    {categories.map( category => (
                                        <label　key={category.id}>
                                            <input type="radio" value={category.name} name="category"
                                                checked={selectedCategory != undefined && selectedCategory.id === category.id}
                                                onChange={handleCategorySelectButton(categories, category.id)}
                                            />{category.name}
                                        </label>
                                    ))}
                                </div>
                            </div> */}
                            <div>
                                <button>Add Tag</button>
                            </div>
                        </form>
                    </div>
                )
            }}
        </Query>
    )
}