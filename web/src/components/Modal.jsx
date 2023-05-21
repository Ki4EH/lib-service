import React from 'react';
import "../styles/modal.css"
const Modal = ({active , setActive, children}) => {
    return (
        <div className={active? "modal active" : " modal"} onClick={() => setActive(false)}>
            <div className="modal__content" onClick={(e) => e.stopPropagation()}>
                {children}
            </div>
        </div>
    );
};

export default Modal;