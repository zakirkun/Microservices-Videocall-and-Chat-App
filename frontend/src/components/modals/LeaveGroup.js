import React, {useContext, useState} from 'react';
import { Modal, ModalHeader, ModalBody } from 'reactstrap';
import { actionTypes, StorageContext } from '../../ChatStorage';
import {DeleteMember} from '../../requests/Groups';

export const ModalLeaveGroup = ({ show, toggle, member, group, setCurrent }) => {

    const [, dispatch] = useContext(StorageContext);

    const [msg, setMsg] = useState("");

    const submit = async() => {
        let response = await DeleteMember(member.groupID, member.ID);
        let flag = false;
        if (response.status !== 200){
            dispatch({type: actionTypes.DELETE_GROUP, payload: group.ID})
            setMsg("You left group");
            flag = true;
        } else {
            setMsg(response.data.err);
        }
        setTimeout(function () {    
            toggle();
            setMsg("");
            if (flag) {
                setCurrent({});
            }
        }, 1000);
    }

    return (
        <Modal id="buy" tabIndex="-1" role="dialog" isOpen={show} toggle={toggle}>
            <div role="document">
                <ModalHeader toggle={toggle} className="bg-dark text-primary text-center">
                    Delete Group
                </ModalHeader>
                <ModalBody>
                    <div>
                        {msg!==""?<h5 className="mb-4 text-danger">{msg}</h5>:null}
                        <div className='form-group'>
                            <label>Are you sure you want to delete group {group.name}?</label>
                        </div>
                        <div className="form-row text-center">
                            <div className="col-12 mt-2">
                                <button className="btn btn-dark btn-large text-primary" onClick={submit}>Delete</button>
                            </div>
                        </div>
                    </div>
                </ModalBody>
            </div>
        </Modal>
    );
} 
