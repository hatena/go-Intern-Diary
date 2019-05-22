import React from "react"

import {Tag} from "./container"

interface AddTagFormProps {
    tagName: string
    handleTagSubmit: (e: React.FormEvent<HTMLFormElement>) => void
    handleTagInput: (e: React.ChangeEvent<HTMLInputElement>) => void
    handleDeleteButton: (selectedTag: Tag) => () => void
    tags: Tag[]
}

export const AddTagForm: React.StatelessComponent<AddTagFormProps> = ({tagName, handleTagSubmit, handleTagInput, handleDeleteButton, tags}) => {
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
                <div>
                    <button>Add Tag</button>
                </div>
            </form>
        </div>
    )
}