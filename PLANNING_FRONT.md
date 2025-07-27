# Frontend Registration Form Requirements

## Registration Form Fields

The landing page registration form needs to collect the following information:

### Required Fields:
1. **Name** - Full name of the participant
2. **Nickname** - Display name/handle
3. **Email** - Contact email address
4. **Region de Vivienda** - Residence region/location
5. **Role** - Select dropdown for participant's preferred role:
   - Frontend Developer
   - Backend Developer
   - Fullstack Developer
   - Other
6. **Languages/Technologies** - Component of tags for multiple choice
7. **Project Idea** - Textarea for describing their project concept
8. **Team Preference** - Radio buttons/select for:
   - Work in teams
   - Work alone
9. **Desired Teammate** - Optional field for specific teammate requests

## Frontend Updates Needed:

1. **Form State Management**: Update useState to include all new fields
2. **Form Validation**: Add validation for email, required fields
3. **UI Components**: 
   - Add email input field
   - Add region select/input
   - Add technologies input (tags/chips)
   - Add team preference radio buttons
   - Update role dropdown with more options
4. **Form Layout**: Reorganize form layout to accommodate new fields
5. **API Integration**: Update form submission to send all fields to backend