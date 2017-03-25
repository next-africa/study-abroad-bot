/**
 * Created by pdiouf on 2017-03-24.
 */
import React from 'react';
import MenuLeft from './ContentHeader';

const styles = {
    sidebar: {
        width: 256,
        height: '100%',
    },
    sidebarLink: {
        display: 'block',
        padding: '16px 0px',
        color: '#757575',
        textDecoration: 'none',
    },
    divider: {
        margin: '8px 0',
        height: 1,
        backgroundColor: '#757575',
    },
    content: {
        padding: '16px',
        height: '100%',
        backgroundColor: 'white',
    },
};

const SidebarContent = (props)=>{
    const style = props.style? {...styles.sidebar, ...props.style}: styles.sidebar;

    const links= [];

    for(let ind = 0; ind < 20; ind++){
        links.push (
            <a key={ind} href="#" style={styles.sidebarLink}> mock menu item {ind}</a>
        )
    }
    return(
        <MenuLeft title="Menu" style={style}>
            <div style={style.content}>
                <a href="index.html" style={style.sidebarLink}>Home</a>
                <div style={styles.divider}/>
                {links}
            </div>
        </MenuLeft>
    );
};

SidebarContent.propTypes = {
    style: React.PropTypes.object
}
export default SidebarContent;