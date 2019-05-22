import React from "react"
import {Tag} from "./container"

interface DiaryTagProps {
    tags: Tag[];
}

export const DiaryTag: React.StatelessComponent<DiaryTagProps> = ({tags}) => {
    return (
        <div className="DiaryTag">
            <h3>この日記のタグ</h3>
            {(tags.length == 0) && <span>タグがありません</span>}
            {tags.map(tag => 
                <span key={tag.name}>{tag.name} </span>)}
        </div>
    )
}