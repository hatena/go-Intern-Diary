import React from "react"
import {Mutation} from "react-apollo"
import {Tag} from "./container"
import { CreateDiary, CreateDiaryVariables } from "../__generated__/CreateDiary"
import { mutation as createDiary } from "../AddDiaryForm/container"
import { createUpdateDiary} from "./container"

interface CreateDiaryFormProps {
    name: string
    handleSubmit: (create: () => void ) => (e: React.FormEvent<HTMLFormElement>) => void
    handleInput: (e: React.ChangeEvent<HTMLInputElement>) => void
    tags: Tag[]
}

export const CreateDiaryForm: React.StatelessComponent<CreateDiaryFormProps> = ({name, handleSubmit, handleInput, tags}) => {
    return (
        <Mutation<CreateDiary, CreateDiaryVariables> mutation={createDiary}>
            {(create) => {
                return (
                    <form className="CreateDiaryForm" onSubmit={handleSubmit(() => create({ variables: {name}, update: createUpdateDiary}))}>
                        <div>
                            <label>Diary Name:
                                <input type="TEXT" name="name" value={name} onChange={handleInput} />
                            </label>
                        </div>
                        <div>
                            <button>Create New</button>
                        </div>
                    </form>
                )
            }}
        </Mutation>
    )
}