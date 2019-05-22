import React from "react"
import {Tag} from "./container"

interface DiaryTagProps {
    tags: Tag[];
}

export const DiaryTag: React.StatelessComponent<DiaryTagProps> = ({tags}) => {
    return (
        <div className="DiaryTag">
            {tags.map(tag => 
                <span key={tag.name}>tag.name </span>)}
        </div>
    )
}