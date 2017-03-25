/**
 * Created by pdiouf on 2017-03-12.
 */
import  React  from  'react'
const Select= (props)=> (

    <div className="form-group">
        name={props.name}
        value={props.selectedOption}
        onChange={props.controlFunc}
        className={"form-select"}
        <option value="">{props.placeholder}</option>
        {props.options.map(opt =>{
            return(
                <option
                    key={opt}
                    value={opt}>{opt} </option>
            );
        })}

    </div>
);

Select.propTypes={
    name: React.PropTypes.string.isRequired,
    options: React.PropTypes.array.isRequired,
    selectedOption: React.PropTypes.string,
    controlFunc:React.PropTypes.func.isRequired,
    placeholder:React.PropTypes.string
};
export default Select;